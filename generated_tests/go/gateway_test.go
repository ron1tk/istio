package core_test

import (
	"testing"

	"istio.io/istio/pilot/pkg/networking/core"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/protocol"
	"istio.io/istio/pkg/config/host"
	"istio.io/istio/pkg/test/mock"
)

func TestMutableGatewayListener_build(t *testing.T) {
	type fields struct {
		Listener *listener.Listener
	}
	type args struct {
		builder *core.ListenerBuilder
		opts    core.gatewayListenerOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Nil HTTP options",
			fields: fields{Listener: mock.NewFakeListener("testListener")},
			args: args{
				builder: mock.NewFakeListenerBuilder(),
				opts: core.gatewayListenerOpts{
					filterChainOpts: []*core.filterChainOpts{
						{
							httpOpts: nil, // Simulating TCP
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:   "Empty filter chains error",
			fields: fields{Listener: mock.NewFakeListener("emptyFilterChainsListener")},
			args: args{
				builder: mock.NewFakeListenerBuilder(),
				opts:    core.gatewayListenerOpts{},
			},
			wantErr: true,
		},
		// Additional test cases for other scenarios...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ml := &core.MutableGatewayListener{
				Listener: tt.fields.Listener,
			}
			if err := ml.build(tt.args.builder, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("MutableGatewayListener.build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigGeneratorImpl_buildGatewayListeners(t *testing.T) {
	// Mocking dependencies and defining test cases for buildGatewayListeners
	// This function likely requires complex setup including mock ConfigGeneratorImpl, ListenerBuilder, and expected Listener configurations
}

func Test_buildGatewayListenerTLSContext(t *testing.T) {
	type args struct {
		mesh         *meshconfig.MeshConfig
		server       *networking.Server
		proxy        *model.Proxy
		protocol     model.Protocol
	}
	tests := []struct {
		name string
		args args
		want *tls.DownstreamTlsContext
	}{
		// Define test cases with different combinations of meshconfig, server TLS settings, and expected DownstreamTlsContext configurations
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.buildGatewayListenerTLSContext(tt.args.mesh, tt.args.server, tt.args.proxy, tt.args.protocol)
			// Use reflect.DeepEqual or a TLS context specific comparison function to compare got and tt.want
		})
	}
}

// Additional tests covering other exported and unexported functions...
```

Note: `mock.NewFakeListener`, `mock.NewFakeListenerBuilder`, and any specific comparison logic for TLS contexts are placeholders indicating where mock objects or specific comparison logic would need to be implemented. 

Testing these components thoroughly may require a significant amount of setup, including creating mock objects that satisfy various interfaces (`ListenerBuilder`, `ConfigGeneratorImpl`, etc.), and potentially using the `gomock` library or similar for generating these mocks. 

Moreover, the testing of some functions, especially those interacting heavily with external dependencies or producing complex configurations, may require integration tests or a more detailed unit test setup to fully cover various edge cases and success/failure scenarios.
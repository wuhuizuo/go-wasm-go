package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuhuizuo/go-wasm-go/runner/plugin"
)

const (
	goPluginSo                = "provider/plugin/ok/plugin.so"
	goPluginSoThird           = "provider/plugin/third/plugin.so"
	goPluginSoThirdDiffModVer = "provider/plugin/third_diff_mod_ver/plugin.so"
	goPluginSo_1_15           = "provider/plugin/ok/plugin-1.15.so"
	goPluginSo_1_16           = "provider/plugin/ok/plugin-1.16.so"
	goPluginSo_1_17_1         = "provider/plugin/ok/plugin-1.17.1.so"
)

func TestPlugin(t *testing.T) {
	t.Run("same go version same thirdparty versions", func(t *testing.T) {
		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSo))
	})

	t.Run("plugin using thirdparty no in main", func(t *testing.T) {
		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSoThird))
	})

	t.Run("plugin using different version of package compared to main", func(t *testing.T) {
		t.Skip("plugin was built with a different version of package ****")

		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSoThirdDiffModVer))
	})

	t.Run("1.17.3 call plugin compiled in 1.17.1", func(t *testing.T) {
		t.Skip("调不了")

		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSo_1_17_1))
	})

	t.Run("1.17 call plugin compiled in 1.16", func(t *testing.T) {
		t.Skip("调不了")

		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSo_1_16))
	})

	t.Run("1.17 call plugin compiled in 1.15", func(t *testing.T) {
		t.Skip("调不了")

		testPlugin(t, filepath.Join(selfDir(t), "..", goPluginSo_1_15))
	})
}

func testPlugin(t *testing.T, pluginSo string) {
	t.Run("algorithm", func(t *testing.T) {
		fbTests := []struct {
			name string
			in   int32
			want int32
		}{
			{name: "5", in: 5, want: 5},
			{name: "10", in: 10, want: 55},
			{name: "20", in: 20, want: 6765},
			{name: "30", in: 30, want: 832040},
		}

		f := plugin.NewGoPluginAlgFn(t, pluginSo, fibFuncName)
		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		f := plugin.NewGoPluginIOFn(t, pluginSo, httpReqFuncName)
		assert.NotNil(t, f)
		f()
	})

	t.Run("file io", func(t *testing.T) {
		f := plugin.NewGoPluginIOErrFn(t, pluginSo, ioFunName)
		if err := f(); err != nil {
			t.Error(err)
		}
	})

	t.Run("multi threads", func(t *testing.T) {
		f := plugin.NewGoPluginMultiThreads(t, pluginSo, multiThreadsFuncName)
		f(4)
	})
}

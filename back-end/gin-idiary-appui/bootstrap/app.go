/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 16:22:52
 * @Description: file content
 */

package bootstrap

import (
	"context"
	"net/http"
	"path/filepath"

	"gin-idiary-appui/httpapi"

	"gin-idiary-appui/library/conf"
	"gin-idiary-appui/library/env"
	"gin-idiary-appui/library/logit"
)

const (
	appConfPath = "./conf/app.toml"
)

// Config app's conf
// default conf/app.toml
type Config struct {
	APPName string
	IDC     string
	RunMode string

	Env env.AppEnv

	// conf of http service
	HTTPServer struct {
		Listen       string
		ReadTimeout  int // ms
		WriteTimeout int // ms
		IdleTimeout  int
	}
}

// ParserAppConfig
func ParserAppConfig(filePath string) (*Config, error) {
	confPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	var c *Config
	if err := conf.Parse(confPath, &c); err != nil {
		return nil, err
	}
	// parse and set global conf
	rootDir := filepath.Dir(filepath.Dir(confPath))
	opt := env.Option{
		AppName: c.APPName,
		RunMode: c.RunMode,
		RootDir: rootDir,
		DataDir: filepath.Join(rootDir, "data"),
		LogDir:  filepath.Join(rootDir, "log"),
		ConfDir: filepath.Join(rootDir, filepath.Base(filepath.Dir(confPath))),
	}
	c.Env = env.New(opt)
	return c, nil
}

// App application
type App struct {
	ctx    context.Context
	config *Config
	close  func()
}

// NewApp establish an APP
func NewApp(ctx context.Context, c *Config) *App {
	ctxRet, cancel := context.WithCancel(ctx)
	app := &App{
		ctx:    ctxRet,
		config: c,
		close:  cancel,
	}
	return app
}

// Start start the service
func (app *App) Start() error {
	// start distribute routers
	httpapi.InitRouters()
	// start record logs
	logit.Init("ziweiapp")
	logit.Logger.Info("APP START")
	// start listening to port
	logit.Logger.Info("APP listening at: %s", app.config.HTTPServer.Listen)
	err := http.ListenAndServe(app.config.HTTPServer.Listen, nil)
	if err != nil {
		return err
	}
	return nil
}

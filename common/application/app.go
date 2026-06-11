package application

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dannyhenry/billing/common/database"
	lsr "github.com/dannyhenry/billing/internal/domain/loan-schedules/repository"
	lss "github.com/dannyhenry/billing/internal/domain/loan-schedules/service"
	lr "github.com/dannyhenry/billing/internal/domain/loan/repository"
	ls "github.com/dannyhenry/billing/internal/domain/loan/service"
	lh "github.com/dannyhenry/billing/internal/handler/loan"
	"github.com/dannyhenry/billing/internal/infrastructure/rest"
	"github.com/labstack/echo/v4"
)

type App struct {
	E         *echo.Echo
	DBManager database.DatabaseManager
}

func (app *App) Start(addr string) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func(sig chan os.Signal) {
		if err := app.E.Start(addr); err != nil {
			sig <- syscall.SIGTERM
		}
	}(sig)

	<-sig
}

func (app *App) Initialize() {
	app.InitializeRoutes()
}

func (app *App) InitializeRoutes() {
	loanScheduleRepository := lsr.NewRepository(app.DBManager)
	loanRepository := lr.NewRepository(app.DBManager)

	loanService := ls.NewService(loanRepository)
	loanScheduleService := lss.NewService(loanScheduleRepository)
	loanHandler := lh.NewHandler(loanService, loanScheduleService)

	rest.RegisterRoutes(app.E, loanHandler)
}

func (app *App) InitializeDatabase(dsn string, maxIdleConns int, maxOpenConns int) {
	err := app.DBManager.Initialize(dsn, maxIdleConns, maxOpenConns)

	if err != nil {
		panic(err)
	}
}

package delivery

import (
	"fmt"

	"github.com/NursiNursi/laundry-apps/config"
	"github.com/NursiNursi/laundry-apps/delivery/controller"
	"github.com/NursiNursi/laundry-apps/manager"
	"github.com/NursiNursi/laundry-apps/utils/exceptions"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine     *gin.Engine
	host       string
}

func (s *Server) Run() {
	s.setupControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupControllers() {
	// semua controller disini
	controller.NewUomController(s.useCaseManager.UomUseCase(), s.engine)
	controller.NewProductController(s.engine, s.useCaseManager.ProductUseCase())
	controller.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
	controller.NewEmployeeController(s.engine, s.useCaseManager.EmployeeUseCase())
	controller.NewBillController(s.engine, s.useCaseManager.BillUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:     engine,
		host:       host,
	}
}

package main

import "github.com/TanguiLouedec/flickhive_back/pkg/logger"

func main()  {
  logger.InitLogger()
  log := logger.GetLogger()
  defer log.Sync()

  log.Info("🚀 Starting Movie App Backend Server")
}

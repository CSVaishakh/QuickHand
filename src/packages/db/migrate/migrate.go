package migrate

import (
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func Up() error {
	_ = godotenv.Load("src/packages/db/.env.local")

	cmd := exec.Command(
		"migrate",
		"-path", "src/packages/db/migrations",
		"-database", os.Getenv("DATABASE_URL"),
		"up",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
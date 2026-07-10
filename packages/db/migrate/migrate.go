package migrate

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func runMigrate(args ...string) error {
	_ = godotenv.Load("/packages/db/.env.local")

	cmdArgs := append([]string{
		"-path", "packages/db/migrations",
		"-database", os.Getenv("DATABASE_URL"),
	}, args...)

	cmd := exec.Command("migrate", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Up() error {
	return runMigrate("up")
}

func Down() error {
	return runMigrate("down")
}

func Version() error {
	return runMigrate("version")
}

func Force(version int) error {
	return runMigrate("force", fmt.Sprintf("%d", version))
}

func Create(name string) error {
	return runMigrate("create", "-ext", "sql", "-dir", "packages/db/migrations", "-seq", name)
}
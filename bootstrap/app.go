package bootstrap

func Initialize() {
	env := NewEnv()
	InitializeDatabase(env)
}

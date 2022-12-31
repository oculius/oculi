package rest

func (w *webServer) BeforeRun(hf HookFunction) Server {
	w.beforeRun = append(w.beforeRun, hf)
	return w
}

func (w *webServer) AfterRun(hf HookFunction) Server {
	w.afterRun = append(w.afterRun, hf)
	return w
}

func (w *webServer) BeforeExit(hf HookFunction) Server {
	w.beforeExit = append(w.beforeExit, hf)
	return w
}

func (w *webServer) AfterExit(hf HookFunction) Server {
	w.afterExit = append(w.afterExit, hf)
	return w
}

func (w *webServer) apply(hooks []HookFunction) error {
	for _, fn := range hooks {
		if err := fn(w.resource); err != nil {
			return err
		}
	}
	return nil
}

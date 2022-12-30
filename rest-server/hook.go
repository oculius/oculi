package rest

func (w *WebServer) BeforeRun(hf HookFunction) Engine {
	w.beforeRun = append(w.beforeRun, hf)
	return w
}

func (w *WebServer) AfterRun(hf HookFunction) Engine {
	w.afterRun = append(w.afterRun, hf)
	return w
}

func (w *WebServer) BeforeExit(hf HookFunction) Engine {
	w.beforeExit = append(w.beforeExit, hf)
	return w
}

func (w *WebServer) AfterExit(hf HookFunction) Engine {
	w.afterExit = append(w.afterExit, hf)
	return w
}

func (w *WebServer) apply(hooks []HookFunction) error {
	for _, fn := range hooks {
		if err := fn(w.resource); err != nil {
			return err
		}
	}
	return nil
}

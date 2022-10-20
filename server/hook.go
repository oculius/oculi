package server

func (w *WebServer) BeforeRun(hf HookFunction) IServer {
	w.beforeRun = append(w.beforeRun, hf)
	return w
}

func (w *WebServer) AfterRun(hf HookFunction) IServer {
	w.afterRun = append(w.afterRun, hf)
	return w
}

func (w *WebServer) BeforeExit(hf HookFunction) IServer {
	w.beforeExit = append(w.beforeExit, hf)
	return w
}

func (w *WebServer) AfterExit(hf HookFunction) IServer {
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

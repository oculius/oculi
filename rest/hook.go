package rest

func (w *webServer) BeforeRun(hf ...HookFunction) {
	if len(hf) == 0 {
		return
	}
	w.beforeRun = append(w.beforeRun, hf...)
}

func (w *webServer) AfterRun(hf ...HookFunction) {
	if len(hf) == 0 {
		return
	}
	w.afterRun = append(w.afterRun, hf...)
}

func (w *webServer) BeforeExit(hf ...HookFunction) {
	if len(hf) == 0 {
		return
	}
	w.beforeExit = append(w.beforeExit, hf...)
}

func (w *webServer) AfterExit(hf ...HookFunction) {
	if len(hf) == 0 {
		return
	}
	w.afterExit = append(w.afterExit, hf...)
}

func (w *webServer) apply(hooks []HookFunction) error {
	for _, fn := range hooks {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

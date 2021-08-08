package bridge

type Chrome interface {
	OpenURL()
}

type Beta struct {
	OperatingSystem OS
}

func (b Beta) OpenURL() {
	b.OperatingSystem.OpenLink()
}

func NewChromeBetaForUbuntu() Chrome {
	return &Beta{
		OperatingSystem: Ubuntu{},
	}
}

func NewChromeBetaForWindows() Chrome {
	return &Beta{
		OperatingSystem: Windows{},
	}
}



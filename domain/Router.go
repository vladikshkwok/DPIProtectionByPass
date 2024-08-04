package domain

type RouterStat struct {
	Name string
	Text string
}

func NewRouterStat(name string, text string) RouterStat {
	return RouterStat{Name: name, Text: text}
}

type RouterInfo struct {
	Stats []RouterStat
}

type PageData struct {
	DpiProp *DPIProp
	Router  *RouterInfo
	Domains []string
}

func NewPageData() *PageData {
	return &PageData{DpiProp: &DPIProp{}, Router: &RouterInfo{}}
}

type DPIProp struct {
	Status string
}

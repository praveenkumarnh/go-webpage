package webpage

// [PageData]
type PageData map[string]interface{}

// [New]

func NewPageData() *PageData { // {{{
	pageData := make(PageData, 0)

	return &pageData
} // }}}

//

func (this *PageData) Has(name string) bool { // {{{
	_, ok := (*this)[name]

	return ok
} // }}}

func (this *PageData) Set(name string, data interface{}) bool { // {{{
	(*this)[name] = data

	return true
} // }}}

func (this *PageData) Get(name string) interface{} { // {{{
	return (*this)[name]
} // }}}

func (this *PageData) Del(name string) bool { // {{{
	delete((*this), name)

	return true
} // }}}

func (this *PageData) Clear() bool { // {{{
	for name, _ := range *this {
		this.Del(name)
	}

	return true
} // }}}

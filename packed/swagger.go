package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBgN/saxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwOj/uHvCW/CNuaJGAq8/Sw6baFN0fIp1xUvNrawqC4yuLfT1vmamfHmLqZdTWwXGTQZOvsm9ZV4uNSce1bHfLN89p2j9/Y6Czp2bZr6/efZ2zc/zp4Y/55DsWUJc+i0hiKdJ/feefCJrdgorH486WBnR9m0qc8YzrE9/7DvYgKnkszFp8u+dC3fomd4p3zpxuwbTa86hE3eiThI9aY/dpqR02dnIt//PN5/1v35TxWWBseGcvEzf3X8E9L2WHQVl4SByneVm/z+szx+e9h9eRivaigVnfbl/AzxoPn1v+p+vUsv/rJP/9gqrVXXWSSu296J8LAsjLgbvObXDA797OaAExsT1ih6z3n6+L+r6iz5/Of37+UEs63v4JMRC49oY6mdtVGAqTKdUyP7lsdN6V4rc9dlGvzvvy/1q+Ka+kq0Z89tjskLyrIjDhyJEL6SeWPZ7t8ZOyvORGRPmxuVKfrphVXRv21W8m0HT0/P+HlZ//jZk+tdv+dk/Pyw1FrKsXBd4b/wx5urGXc6lhzLv1u3/dlK589eR0/0TpvfpMWgn/tj9evfszkrjBnDOa9XMR0LrRaN8kz963RlariRXvTqnennHT+3Onn2n2Pe9NKCb2Us4132q1nSpwX0ml76p/EEsKzOnRKzbPfUJFaXVXeUzqfu/vs28s2ZfywMDP//B3izczSe8fzxgJGB4SMXAwMsqTAwcKElFXZEUgGnDpPD3xNAupHVBHgzMokwI5IassmgpAYD2xpBJN6EhzAKu1MgQIDhv6M2EwMWh7GygeSZGJgYOhkYGDKZQDxAAAAA//8NTHUcCAMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}

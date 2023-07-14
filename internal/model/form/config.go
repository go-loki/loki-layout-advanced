package form

type IndexV1Req struct {
	Age  int    `form:"age" vd:"$>0"`
	Name string `form:"name" vd:"len($)>3"`
}

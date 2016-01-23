package ri

/* Polygon nvertices is the number of vertices in a single closed planar convex polygen */
func (ctx *Context) Polygon(nvertices RtInt, parameterlist ...Rter) error {

	/* NOTE: we don't need nvertices in the RIB output */
	return ctx.writef("Polygon", parameterlist...)
}

/* GeneralPolygon Define a general planar concave polygon with holes */
func (ctx *Context) GeneralPolygon(nloops RtInt, nvertices RtIntArray, parameterlist ...Rter) error {

	/* NOTE: we don't need nloops in the RIB output */
	var out = []Rter{nvertices}
	out = append(out, parameterlist...)
	return ctx.writef("GeneralPolygon", out...)
}

/* PointsPolygons Degine npolys planar convex polygons that share vertices */
func (ctx *Context) PointsPolygons(npolys RtInt, nvertices RtIntArray, vertices RtIntArray, parameterlist ...Rter) error {

	/* NOTE: we don't need npolys in the RIB output */
	var out = []Rter{nvertices, vertices}
	out = append(out, parameterlist...)
	return ctx.writef("PointsPolygon", out...)
}

/* PointsGeneralPolygons */
func (ctx *Context) PointsGeneralPolygons(nploys RtInt, nloops, nvertices, vertices RtIntArray, parameterlist ...Rter) error {

	var out = []Rter{nloops, nvertices, vertices}
	out = append(out, parameterlist...)
	return ctx.writef("PointsGeneralPolygons", out...)
}

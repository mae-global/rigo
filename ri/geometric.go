package ri

/* Polygon nvertices is the number of vertices in a single closed planar convex polygen */
func (r *Ri) Polygon(nvertices RtInt, parameterlist ...Rter) error {

	/* NOTE: we don't need nvertices in the RIB output */
	var out = []Rter{PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("Polygon", out...)
}

func (r *Ri) PolygonV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Polygon", out...)
}

/* GeneralPolygon Define a general planar concave polygon with holes */
func (r *Ri) GeneralPolygon(nloops RtInt, nvertices RtIntArray, parameterlist ...Rter) error {

	/* NOTE: we don't need nloops in the RIB output */
	var out = []Rter{nvertices, PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("GeneralPolygon", out...)
}

func (r *Ri) GeneralPolygonV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("GeneralPolygon", out...)
}

/* PointsPolygons Degine npolys planar convex polygons that share vertices */
func (r *Ri) PointsPolygons(npolys RtInt, nvertices RtIntArray, vertices RtIntArray, parameterlist ...Rter) error {

	/* NOTE: we don't need npolys in the RIB output */
	var out = []Rter{nvertices, vertices, PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("PointsPolygons", out...)
}

func (r *Ri) PointsPolygonsV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("PointsPolygons", out...)
}

/* PointsGeneralPolygons */
func (r *Ri) PointsGeneralPolygons(nploys RtInt, nloops, nvertices, vertices RtIntArray, parameterlist ...Rter) error {

	var out = []Rter{nloops, nvertices, vertices, PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("PointsGeneralPolygons", out...)
}

func (r *Ri) PointsGeneralPolygonsV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("PointsGeneralPolygons", out...)
}

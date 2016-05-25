package math

// CartesianToSpherical converts 3-dimensional cartesian coordinates (x,y,z) to
// spherical coordinates with radius r, inclination theta, and azimuth phi.
func CartesianToSpherical(coord Vec3) (r, theta, phi float32) {
	r = coord.Len()
	theta = Acos(coord[2] / r)
	phi = Atan2(coord[1], coord[0])
	return
}

// SphericalToCartesian converts spherical coordinates with radius r,
// inclination theta, and azimuth phi to cartesian coordinates (x,y,z).
func SphericalToCartesian(r, theta, phi float32) Vec3 {
	st, ct := Sincos(theta)
	sp, cp := Sincos(phi)

	return Vec3{r * st * cp, r * st * sp, r * ct}
}

// CartesianToCylindrical converts 3-dimensional cartesian coordinates (x,y,z)
// to cylindrical coordinates with radial distance r, azimuth phi, and height z.
func CartesianToCylindrical(coord Vec3) (rho, phi, z float32) {
	rho = Hypot(coord[0], coord[1])
	phi = Atan2(coord[1], coord[0])
	z = coord[2]
	return
}

// CylindricalToCartesian converts cylindrical coordinates with radial distance
// r, azimuth phi, and height z to cartesian coordinates (x,y,z).
func CylindricalToCartesian(rho, phi, z float32) Vec3 {
	s, c := Sincos(phi)

	return Vec3{rho * c, rho * s, z}
}

// SphericalToCylindrical converts spherical coordinates with radius r,
// inclination theta, and azimuth phi to cylindrical coordinates with radial
// distance r, azimuth phi, and height z.
func SphericalToCylindrical(r, theta, phi float32) (rho, phi2, z float32) {
	s, c := Sincos(theta)

	rho = r * s
	z = r * c
	phi2 = phi

	return
}

// CylindricalToSpherical converts cylindrical coordinates with radial distance
// r, azimuth phi, and height z to spherical coordinates with radius r,
// inclination theta, and azimuth phi.
func CylindricalToSpherical(rho, phi, z float32) (r, theta, phi2 float32) {
	r = Hypot(rho, z)
	phi2 = phi
	theta = Atan2(rho, z)
	return
}

// DegToRad converts degrees to radians
func DegToRad(angle float32) float32 {
	return angle * Pi / 180
}

// RadToDeg converts radians to degrees
func RadToDeg(angle float32) float32 {
	return angle * 180 / Pi
}

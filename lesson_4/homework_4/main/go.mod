module homework_4

go 1.22.5

replace greeting => ../greeting

replace weather => ../weather

replace trafficlight => ../traffic

replace grade => ../grade

replace discount => ../discount

replace temperature => ../temperature

require (
	discount v0.0.0-00010101000000-000000000000
	grade v0.0.0-00010101000000-000000000000
	greeting v0.0.0-00010101000000-000000000000
	temperature v0.0.0-00010101000000-000000000000
	trafficlight v0.0.0-00010101000000-000000000000
	weather v0.0.0-00010101000000-000000000000
)

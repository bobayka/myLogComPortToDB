package domains

const (
	ax = 0
	ay = 1
	az = 2
	gx = 3
	gy = 4
	gz = 5
)

type GyroData struct {
	X, Y, Z string
}

type AccelData struct {
	X, Y, Z string
}

type GyroAccelData struct {
	AccelData
	GyroData
}

func NewGyroAccelData(mas []string) *GyroAccelData {
	return &GyroAccelData{
		AccelData: AccelData{
			X: mas[ax],
			Y: mas[ay],
			Z: mas[az],
		},
		GyroData: GyroData{
			X: mas[gx],
			Y: mas[gy],
			Z: mas[gz],
		}}
}

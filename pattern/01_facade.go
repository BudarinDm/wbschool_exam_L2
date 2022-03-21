package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

//методы которые мы не можем изменить и упростить
//у каждой из этих струтур свои методы и тд и тп
type VideoFile struct {
}

type OggCompressionCodec struct {
}

type MPEG4CompressionCodec struct {
}

type CodecFactory struct {
	file VideoFile
}

type BitrateReader struct {
	name string
}

type AudioMixer struct {
	buffer BitrateReader
}

//наш фасад
type VideoConverter struct {
}

//фасад не имеет всей функциональности фреймворка, но зато скрывает его сложность от клиентов.
//наши методы фасада
func (v *VideoConverter) Convert(name, format string) string {
	file := new(VideoFile)
	_ = CodecFactory{
		file: *file,
	}
	if format == "mp4" {
		//destinationCodec := new(MPEG4CompressionCodec)
		//типо сложная логика
	}

	buffer := BitrateReader{name: name}
	result := AudioMixer{
		buffer: buffer,
	}
	fmt.Println(result.buffer.name)
	return "типо сконвертированый файл"
}

func main() { //по сути пользователи вызывают только 1 метод вместо нескольких
	convertor := new(VideoConverter) //и создают только 1 обьект у себя
	mp4 := convertor.Convert("name", "mp4")

	fmt.Println(mp4)
}

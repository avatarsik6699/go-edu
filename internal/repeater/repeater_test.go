package repeater

import "testing"

func TestRepeater(t *testing.T) {
	t.Run("should return seq aaa", func(t *testing.T) {
		got := Repeater("a", 3)
		want := "aaa"

		if got != want {
			t.Errorf("want -> %s, got -> %s", want, got)
		}
	})

	t.Run("should return empty string if count 0", func(t *testing.T) {
		got := Repeater("a", 0)
		want := ""

		if got != want {
			t.Errorf("want -> %s, got -> %s", want, got)
		}
	})
}

func BenchmarkRepeater(b *testing.B) {
	// b.N - это количество итераций, которое Go автоматически подберет
	// для получения статистически значимых результатов
	for i := 0; i < b.N; i++ {
		Repeater("a", 10)
	}
}

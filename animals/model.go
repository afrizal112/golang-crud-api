package animals

type Animal struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Species   string `json:"species"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

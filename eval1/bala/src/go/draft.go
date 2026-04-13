/*
## Contexto
Um atirador de elite estava atirando a 400 metros de distância em um papel milimetrado. Para calcular a imprecisão da arma,
 ele dava dois tiros e media a distância entre eles. Como ele não tinha uma régua, ele pegou as coordenadas dos pontos
  no plano cartesiano e usou a fórmula da distância entre dois pontos.

$$d_{AB} = \sqrt{(x₂ - x₁)^2 + (y₂ - y₁)^2}$$

Dada a fórmula da distância entre dois pontos e os valores x e y de cada ponto,
imprima a distância entre os pontos com duas casas decimais.
### Entrada

- Coordenada ***X*** e coordenada ***Y*** do primeiro ponto.
- Coordenada ***X*** e coordenada ***Y*** do segundo ponto.

### Saída
- A distância entre os pontos com duas casas decimais.
*/

package main
import(
    "fmt"
    "math"
)
func main() {
    var a1, a2, b1, b2, form float64

    fmt.Scan(&a1, &b1)
    fmt.Scan(&a2, &b2)

    form = math.Sqrt((a1-a2)*(a1-a2)+(b1-b2)*(b1-b2))

    fmt.Printf("%.2f\n", form)
}

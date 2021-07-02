//queue follows the FIFO pattern

package main 
import "fmt"


func enqueue (queue[] string, element string ) [] string {
	queue = append(queue, element)
	fmt.Println(queue)
	return queue

}

func dequeue(queue[] string) [] string{
	firstElement := queue[0]
	fmt.Println("Removed", firstElement)
	queue = queue[1:]
	fmt.Println("New queue", queue)
	return queue
}

func main() {
	var queue[] string
	fmt.Println("enqueing")
	queue = enqueue(queue, "apple")
	queue = enqueue(queue, "bag")
	queue = enqueue(queue, "carrot")
	queue = enqueue(queue, "dog")

	fmt.Println("dequeing")
	queue = dequeue(queue)



}
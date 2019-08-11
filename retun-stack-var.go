package main
import (
"fmt");


type abc struct {
  a int;
  b int;
  c string;
};

func return_stack_var() *abc {
 return &abc{10, 100, "Hi "};
}

func main() {
   
 a := return_stack_var();
 b := func () *abc {
return &abc{10, 100, "Hi "};
} ();

 fmt.Println(" a.a ", a.a, b.a);
}
package builder

import "fmt"

/**

Copyright 2022, Vinicius Nordi <viniciusnordiesperanca@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

**/

type Dep1 struct{}

func (d *Dep1) String() string {
	return "I am the db"
}

type Dep2 struct{}

func (d *Dep2) String() string {
	return "I am an API"
}

type Dep3 struct{}

func (d *Dep3) String() string {
	return "I am a message queue"
}

type Service interface {
	ActionDep1() error
	ActionDep2() error
	ActionDep3() error
}

type serviceImpl struct {
	dep1 *Dep1
	dep2 *Dep2
	dep3 *Dep3
}

func (s *serviceImpl) ActionDep1() error {
	fmt.Println("Printing ActionDep1")
	fmt.Println(s.dep1)
	fmt.Println("Finishing ActionDep1")
	return nil
}

func (s *serviceImpl) ActionDep2() error {
	fmt.Println("Printing ActionDep2")
	fmt.Println(s.dep2)
	fmt.Println("Finishing ActionDep2")
	return nil
}

func (s *serviceImpl) ActionDep3() error {
	fmt.Println("Printing ActionDep3")
	fmt.Println(s.dep3)
	fmt.Println("Finishing ActionDep3")
	return nil
}

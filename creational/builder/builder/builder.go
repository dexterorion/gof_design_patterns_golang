package builder

/**

Copyright 2022, Vinicius Nordi <viniciusnordiesperanca@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

**/

type ServiceBuilder interface {
	SetDependency1(d *Dep1) ServiceBuilder
	SetDependency2(d *Dep2) ServiceBuilder
	SetDependency3(d *Dep3) ServiceBuilder
	Build() Service
}

type serviceBuilderImpl struct {
	dep1 *Dep1
	dep2 *Dep2
	dep3 *Dep3
}

func NewServiceBuilder() ServiceBuilder {
	return &serviceBuilderImpl{}
}

func (sb *serviceBuilderImpl) SetDependency1(d *Dep1) ServiceBuilder {
	sb.dep1 = d
	return sb
}

func (sb *serviceBuilderImpl) SetDependency2(d *Dep2) ServiceBuilder {
	sb.dep2 = d
	return sb
}

func (sb *serviceBuilderImpl) SetDependency3(d *Dep3) ServiceBuilder {
	sb.dep3 = d
	return sb
}

func (sb *serviceBuilderImpl) Build() Service {
	return &serviceImpl{
		dep1: sb.dep1,
		dep2: sb.dep2,
		dep3: sb.dep3,
	}
}

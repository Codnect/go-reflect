package reflect

type Package struct {
	name string
}

func newPackage(packageName string) *Package {
	return &Package{
		name: packageName,
	}
}

func (p *Package) GetName() string {
	return p.name
}

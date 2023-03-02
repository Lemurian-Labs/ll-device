package main

int main() {
	v1 := make([]float32, 1024)
	for i := range v1 {
		v1[i] = rand.NormFloat32()
	}
	printf("vector v1: %v\n", v1);
	return 0;
}

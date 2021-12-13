function fibonacci(param) {
    if (param <= 1) {
        return param;
    } 

	return fibonacci(param - 1) + fibonacci(param - 2);
}

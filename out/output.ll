define i32 @main() {
entry:
	%0 = alloca i32
	store i32 0, i32* %0
	%1 = alloca i32
	store i32 3, i32* %1
	%2 = load i32, i32* %1
	%3 = mul i32 %2, 2
	%4 = sub i32 %3, 4
	%5 = add i32 %4, 5
	store i32 %5, i32* %0
	%6 = load i32, i32* %0
	%7 = add i32 2, 1
	%8 = srem i32 45, %7
	%9 = add i32 %6, %8
	store i32 %9, i32* %0
	%10 = load i32, i32* %0
	ret i32 %10
}

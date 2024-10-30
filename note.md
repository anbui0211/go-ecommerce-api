# Note

1. **Context**: dùng để kiểm soát luồng đi của 1 vòng đời reuqest

2. **Pool connect trong MySQL**: Mở sẵn các kết nối để khi nào user cần thì có thể lấy để sử dụng

3. **Benchmark**: Improving function performance
Example: Performance benchmark comparing MySQL with 1 connection pools to 10 connections pools
    | Benchmark Name | Number of Operations | Average Time per Operation |Average Bytes Allocated per Operation | Average Allocations per Operation |
    | :-------- | :------- | :-------- |:-------- |:-------- |
    | `BenchmarkMaxOpenConns1-8` | `1299` | `977401 ns/op` |`5752 B/op` |`75 allocs/op` |
    | `BenchmarkMaxOpenConns10-8` | `4789` | `221251 ns/op` |`5674 B/op` |`75 5674/op` |

4. Wire Dependency Injection
    - Inversion of Control (IoC): Là nguyên tắc thiết kế trong OOP
    - Dependency Injection (DI): là phương pháp phổ biến trong Inversion of Control

5. **Wire**
    Mục đích: Giúp tự động hóa quá trình tạo ra các đối tượng có nhiều dependency (phụ thuộc).

    ``` go
      // 'cd' to wire folder and run cmd
      wire
    ```

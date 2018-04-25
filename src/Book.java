public class Book {
    String name;
    int num = 1;

    public Book(String name, int num) {
        this.name = name;
        this.num = num;
    }

    public Boolean addNum(int addition) {
        num += addition;

        return true;
    }

    @Override
    public boolean equals(Object object) {
        if(object instanceof Book){
            Book book = (Book)object;
            return this.name.equals(book.name);
        }
        return false;
    }
}
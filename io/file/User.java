package like.entity;

public class User {

  private String userName;

  private int age;

  private boolean isMan;

  public User() {
  }

  public User(String userName, int age, boolean isMan) {
    this.userName = userName;
    this.age = age;
    this.isMan = isMan;
  }

  public String getUserName() {
    return userName;
  }

  public void setUserName(String userName) {
    this.userName = userName;
  }

  public int getAge() {
    return age;
  }

  public void setAge(int age) {
    this.age = age;
  }

  public boolean isMan() {
    return isMan;
  }

  public void setMan(boolean man) {
    isMan = man;
  }
}
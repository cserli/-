#include <windows.h>
#include <conio.h>
// 使用了WinAPI来移动控制台的光标
void gotoxy(int x,int y)
{
	HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE);  // 获取控制台句柄
    SetConsoleTextAttribute(handle, FOREGROUND_INTENSITY | FOREGROUND_RED); // 设置为红色
    COORD c;
    c.X=x,c.Y=y;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}
// 从键盘获取一次按键，但不显示到控制台
int direct()
{
    return _getch();
}
// 控制台清屏数据
void system_cls()
{
	system("cls");
}
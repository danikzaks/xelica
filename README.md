# Xelica Tool for Developers

Xelica — это многофункциональный инструмент для разработчиков, который предоставляет простой и эффективный интерфейс для форматирования JSON, перевода чисел между различными системами счисления и других полезных функций. Приложение разработано с использованием фреймворка Fyne, что позволяет создавать кросс-платформенные графические приложения.

## Функции:
- Форматирование JSON: Легко форматируйте JSON-данные для улучшения читаемости.

- Перевод чисел между системами счисления: Преобразование чисел между различными системами счисления (например, из десятичной в двоичную, шестнадцатеричную и наоборот).

- Поддержка других инструментов для разработчиков: В дальнейшем планируется добавление новых полезных инструментов для разработчиков.

## Установка:
### Для Windows, macOS и Linux

1. Скачайте последнюю версию релиза с [GitHub Releases](https://github.com/danikzaks/xelica).

2. Распакуйте архив в удобное место.

3. Запустите приложение двойным кликом на xelica.

## Сборка из исходников

1. Убедитесь, что у вас установлен Go (версия 1.18 и выше).
2. Клонируйте репозиторий:

```bash
git clone https://github.com/danikzaks/xelica.git
cd xelica
```

3. Соберите приложение:

```bash
make build
```

4. Запустите приложение:
```bash
make run
```

## Использование
Запустив приложение, вы получите простой графический интерфейс с несколькими основными функциями:
1. Форматирование JSON: Введите или вставьте ваш JSON в текстовое поле и нажмите кнопку "Форматировать".
2. Перевод чисел: Введите число в одном из полей и выберите систему счисления для конверсии (десятичная, двоичная, шестнадцатеричная).

## Часто задаваемые вопросы (FAQ)
### Как запустить приложение после установки?
После того как вы скачаете и установите приложение, просто откройте папку с файлом и запустите xelica (или соответствующий исполнимый файл для вашей ОС).
### Как добавить поддержку новых систем счисления?
Если вы хотите добавить поддержку новых систем счисления, просто откройте исходный код в папке cmd и обновите соответствующую логику. Мы приветствуем pull requests для новых функций.
### Почему приложение не запускается на macOS?
Убедитесь, что ваше окружение поддерживает запуск графических приложений с использованием Fyne. Если проблема сохраняется, сообщите об этом в разделе Issues на GitHub.
### Где можно найти исходный код проекта?
Исходный код доступен в репозитории на [GitHub](https://github.com/danikzaks/xelica).

## Вклад
Если вы хотите внести свой вклад в проект, не стесняйтесь создавать issues и pull requests на GitHub. Мы приветствуем участие в разработке!
1. Форкните репозиторий.
2. Создайте новую ветку для ваших изменений.
3. Отправьте pull request с вашими изменениями.

## Лицензия
Этот проект лицензируется под лицензией MIT. См. LICENSE для подробностей.

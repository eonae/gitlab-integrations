# Требования

1. Получаем список репозиториев из указанной группы в gitlab
2. Фильтруем список по маске (regexp - получаем на вход).
3. Сверили содержимое веток A и B (получаем на вход).
4. Если есть отличия - открываем MR из A в B.
5. Отслеживаем успешность всех MR. Деплой запустится сам
6. Отслеживаем состояние пайплайнов деплоя.
7. Выводим инфу в sdtout


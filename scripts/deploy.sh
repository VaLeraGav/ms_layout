#!/bin/bash

PROJECT_NAME="ms_layout"
PWD=$(pwd)
DIR_PWD=$PWD

SERVICE_NAME="$PROJECT_NAME.service"
EXEC_START="$DIR_PWD/build/package/$PROJECT_NAME -config $DIR_PWD/configs/local.yaml"

# Копирование файла службы
if sudo cp "$DIR_PWD/init/$SERVICE_NAME" /etc/systemd/system/; then
    echo "Служба скопирована успешно"
else
    echo "Ошибка при копировании службы" >&2
    exit 1
fi

# Проверяем наличие файла
if [ -f "/etc/systemd/system/$SERVICE_NAME" ]; then
    sudo sed -i "s|EXEC_START|$EXEC_START|" /etc/systemd/system/$SERVICE_NAME
    sudo sed -i "s|PWD|$PWD|" /etc/systemd/system/$SERVICE_NAME
    echo "Перезаписали ExecStart"
else
    echo "Ошибка при измении ExecStart" >&2
    exit 1
fi


# Перезагрузка systemd
if sudo systemctl daemon-reload; then
    echo "systemd перезагружен успешно"
else
    echo "Ошибка при перезагрузке systemd" >&2
    exit 1
fi

# Запуск службы
if sudo systemctl start "$SERVICE_NAME"; then
    echo "Служба запущена успешно"
else
    echo "Ошибка при запуске службы" >&2
    exit 1
fi

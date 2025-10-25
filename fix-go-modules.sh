#!/bin/bash

# fix-modules-manual.sh
# Создаёт go.mod вручную для всех примеров в Go-Design-Patterns
# Работает на macOS и Linux

set -e

echo "🔧 Исправление модулей в Go-Design-Patterns..."

# Chapter01
cd Chapter01/libraries
go mod init libraries
cd ../..

# Chapter02
cd Chapter02
go mod init chapter02
cd ..

# Chapter05/command
cd Chapter05/command/example2
go mod init command-example2
cd ../../..

# Chapter05/strategy — создаём ОДИН модуль на всю папку example2
cd Chapter05/strategy/example2
go mod init strategy-example2
cd ../../..

# Chapter10
cd Chapter10/Memory
go mod init memory
cd ../..

# Chapter08 и Chapter09 — отдельные .go файлы в корне главы
# Но они не содержат package main → пропускаем
# (если нужно — можно добавить, но маловероятно)

echo "✅ Все модули созданы!"
echo ""
echo "📌 ВАЖНО: в Chapter05/strategy/example2/cli/main.go"
echo "   замените импорты с '../...' на:"
echo "      import \"strategy-example2\""
echo "      import \"strategy-example2/shapes\""
echo ""
echo "✅ Теперь вы можете запускать примеры:"
echo "   cd Chapter01/libraries && go run main.go"

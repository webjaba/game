# Структура проекта

```
game/
├── assets/                 # Ресурсы
│   ├── sprites/            # Спрайты
│   ├── shaders/            # GLSL шейдеры
│   ├── maps/               # Tiled карты
│   └── audio/              # Звуки
├── internal/
│   ├── game/               # Основная логика
│   │   ├── settings/       # Общие настройки игры
│   │   │   └── settings.go # Файл с конфигурациями и константами
│   │   ├── scenes/         # Сцены игры
│   │   │   ├── game.go     # Игровая сцена
│   │   │   ├── menu.go     # Меню
│   │   │   └── pause.go    # Пауза
│   │   ├── entities/       # Сущности
│   │   │   ├── player.go   # Entity игрока
│   │   │   ├── object.go
│   │   │   └── type.go     # Описание некоторых общих интерфейсов
│   │   ├── systems/        # Системы (ECS)
│   │   │   ├── render.go
│   │   │   ├── movement.go # Система для движения
│   │   │   └── type.go     # Описание некоторых общих интерфейсов
│   │   ├── world/          # Игровой мир
│   │   │   ├── tilemap.go
│   │   │   └── camera.go
│   │   └── utils/          # Утилиты
│   │       ├── math.go     # Утилиты для вычислений
│   │       ├── object.go   # Описание типов данных
│   │       └── assets.go
│   └── pkg/
│       ├── ecs/            # ECS ядро
│       └── shaders/        # Менеджер шейдеров
├── cmd/
│   └── main/               # Точка входа
│       └── main.go
└── go.mod
```

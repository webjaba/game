# Структура проекта

```
mygame/
├── assets/               # Ресурсы
│   ├── sprites/          # Спрайты
│   ├── shaders/          # GLSL шейдеры
│   ├── maps/             # Tiled карты
│   └── audio/            # Звуки
├── internal/
│   ├── game/             # Основная логика
│   │   ├── scenes/       # Сцены игры
│   │   │   ├── game.go   # Игровая сцена
│   │   │   ├── menu.go   # Меню
│   │   │   └── pause.go  # Пауза
│   │   ├── entities/     # Сущности
│   │   │   ├── player.go
│   │   │   ├── enemy.go
│   │   │   └── object.go
│   │   ├── systems/      # Системы (ECS)
│   │   │   ├── render.go
│   │   │   ├── physics.go
│   │   │   └── ai.go
│   │   ├── world/        # Игровой мир
│   │   │   ├── tilemap.go
│   │   │   └── camera.go
│   │   └── utils/        # Утилиты
│   │       ├── math.go
│   │       └── assets.go
│   └── pkg/
│       ├── ecs/          # ECS ядро
│       └── shaders/      # Менеджер шейдеров
├── cmd/
│   └── main/             # Точка входа
│       └── main.go
└── go.mod
```
# 🎮 Game Engines and Technologies

## 📑 Contents
1. [What is a Game Engine?](#what-is-a-game-engine)
2. [Popular Game Engines](#popular-game-engines)
3. [Technologies in Games](#technologies-in-games)
4. [Choosing a Game Engine](#choosing-a-game-engine)
5. [Engine Comparison](#engine-comparison)

---

## 1. 🤖 What is a Game Engine?

A **Game Engine** is a software platform designed for developing and running video games. It provides ready-made components for physics, graphics, sound, input, and other aspects.

### 🔄 Analogy
Think of a game engine as a **construction kit** for a house:
*   You have ready-made walls, windows, doors
*   You only need to plan the layout and add finishing touches
*   You don't need to reinvent every element

### 🧩 Core Components:
*   **Graphics Renderer** — renders 2D/3D
*   **Physics Engine** — simulates physics
*   **Audio System** — sound playback
*   **Input System** — handles keyboard, mouse, gamepad
*   **Animation System** — character movement
*   **Lighting System** — shadows, light
*   **Particle System** — effects (smoke, fire)

---

## 2. 🏆 Popular Game Engines

### 🎮 Unity
*   **Type:** Commercial (with free version)
*   **Programming Language:** C#
*   **Platforms:** Almost all (PC, mobile, VR/AR, consoles)
*   **Advantages:**
    *   Easy to learn
    *   Large community
    *   Many ready-made assets
*   **Example Games:** Cuphead, Hollow Knight, Pokémon GO
*   **License:** Free up to certain revenue, then percentage of profits

### 🎯 Unreal Engine
*   **Type:** Commercial (with free version)
*   **Programming Language:** C++, Blueprint (visual programming)
*   **Platforms:** PC, consoles, mobile, VR/AR
*   **Advantages:**
    *   High-quality graphics
    *   Built-in editor
    *   Photorealistic rendering
*   **Example Games:** Gears of War, Batman: Arkham Series, **Fortnite**
*   **License:** Free, 5% of revenue after $1M annual revenue

**Example from Fortnite:**
Fortnite uses Unreal Engine for:
- Powerful graphics component with bright, comic-style visuals
- Real-time building system
- Support for large game worlds with hundreds of players
- Regular updates and implementation of new features

### 🧊 Godot
*   **Type:** Open Source
*   **Programming Language:** GDScript (similar to Python), C#, C++, Visual Script
*   **Platforms:** PC, mobile, HTML5, consoles
*   **Advantages:**
    *   Completely free
    *   Open source
    *   Lightweight and fast
    *   2D and 3D support
*   **Example Games:** Katana ZERO, Moonlighter, A Short Hike
*   **License:** MIT (completely free)

### 🌟 GameMaker Studio
*   **Type:** Commercial
*   **Programming Language:** GML (GameMaker Language), visual scripting
*   **Platforms:** PC, mobile, HTML5, consoles
*   **Advantages:**
    *   Excellent for 2D games
    *   Easy start
    *   Powerful 2D tools
*   **Example Games:** Undertale, Hotline Miami, Dead Cells
*   **License:** Paid license for publishing

### 🧱 CryEngine
*   **Type:** Commercial
*   **Programming Language:** C++, Lua, visual programming
*   **Platforms:** PC, consoles
*   **Advantages:**
    *   Photorealistic graphics
    *   Advanced graphic capabilities
    *   Tools for open worlds
*   **Example Games:** Far Cry series, Ryse: Son of Rome
*   **License:** Free with percentage of profits

---

## 3. 🔧 Technologies in Games

### 🎨 Graphics Technologies:
*   **Rasterization** — converting 3D to 2D for display
*   **Ray Tracing** — physically accurate light modeling
*   **Shaders** — programs for graphics processing (Vertex, Fragment, Geometry)
*   **PBR (Physically Based Rendering)** — realistic material calculation
*   **LOD (Level of Detail)** — changing detail based on distance

### ⚡ Optimization:
*   **Occlusion Culling** — don't render hidden objects
*   **Frustum Culling** — don't render objects outside view
*   **Multi-threading** — distributing tasks across CPU cores
*   **GPU Compute** — using GPU for calculations

**Example from Fortnite:**
In Fortnite, optimization technologies are used for:
- **Occlusion Culling** — not rendering buildings hidden behind other objects
- **Frustum Culling** — displaying only objects that the player sees
- **Multi-threading** — distributing tasks across CPUs to support 100 players
- **GPU Compute** — optimizing rendering for different graphics adapters

### 🧠 Artificial Intelligence:
*   **Behavior Trees** — behavior trees for AI
*   **Finite State Machines** — finite automata
*   **Pathfinding (A*, NavMesh)** — path finding
*   **Machine Learning** — AI learning based on data

**Example from Fortnite:**
In Fortnite, AI is used for:
- **Finite State Machines** — bot behavior in various modes
- **Pathfinding** — bots moving through the game world
- **Behavior Trees** — decision-making by bots (build, attack, evade)

### 🌐 Networking Technologies:
*   **Client-Server** — client-server architecture
*   **Peer-to-Peer** — peer-to-peer network
*   **Authoritative Server** — server as single source of truth
*   **Lag Compensation** — delay compensation

**Example from Fortnite:**
In Fortnite, the following networking technologies are used:
- **Client-Server** — centralized architecture to support 100 players
- **Authoritative Server** — server as single source of truth to prevent cheating
- **Lag Compensation** — delay compensation for fair play

---

## 4. 🧭 Choosing a Game Engine

### 🎯 Selection Factors:

#### 📱 Target Platform:
*   **Mobile:** Unity, GameMaker, Godot
*   **PC/Consoles:** Unreal Engine, Unity
*   **VR/AR:** Unity, Unreal Engine
*   **2D:** Godot, GameMaker, Unity
*   **3D:** Unreal Engine, Unity, CryEngine

#### 👨‍💻 Skill Level:
*   **Beginners:** GameMaker, Godot, Unity
*   **Advanced:** Unreal Engine, CryEngine
*   **Professionals:** Any, depending on project

#### 💰 Budget:
*   **Budget Projects:** Godot (free)
*   **Individual Developers:** Unity (up to income limit), Godot
*   **Commercial Studios:** Unreal Engine, Unity, CryEngine

#### 🏢 Team Size:
*   **Single Developer:** Godot, GameMaker
*   **Small Team:** Unity, Godot
*   **Large Studio:** Unreal Engine, CryEngine

---

## 5. 📊 Engine Comparison

| Criterion | Unity | Unreal Engine | Godot | GameMaker |
|----------|-------|---------------|-------|-----------|
| Programming Language | C# | C++, Blueprint | GDScript, C# | GML, C++ |
| 2D Support | Good | Average | Excellent | Excellent |
| 3D Support | Excellent | Excellent | Good | Average |
| Graphics | Good | Excellent | Good | Average |
| Community | Large | Large | Growing | Medium |
| Difficulty | Medium | Medium | Low | Low |
| License | Revenue model | Revenue model | MIT | Paid |
| Platforms | All | Almost all | All | Most |

### 🎯 Recommendations:
*   **For Beginners:** Godot or GameMaker
*   **For 2D Games:** Godot or GameMaker
*   **For 3D AAA:** Unreal Engine
*   **For Cross-Platform:** Unity
*   **For Indie Projects:** Godot
*   **For VR/AR:** Unity or Unreal Engine

---

## 6. 🚀 Conclusion

Choosing a game engine is one of the key decisions in game development. The right choice depends on many factors: game type, target platform, development team, and budget. Modern engines provide powerful tools that allow even small teams to create impressive games.

The next materials will cover game programming, graphics technologies, and other development aspects.

<!-- QUIZ_START
[
    {
        "question": "Which programming language is used in the Unity engine?",
        "options": [
            "C++",
            "Java",
            "C#",
            "Python"
        ],
        "correctIndex": 2
    },
    {
        "question": "Which game engine is completely open source?",
        "options": [
            "Unity",
            "Unreal Engine",
            "Godot",
            "GameMaker Studio"
        ],
        "correctIndex": 2
    },
    {
        "question": "What does the abbreviation PBR mean in the context of game graphics?",
        "options": [
            "Primary Behavior Recognition",
            "Physical Based Rendering",
            "Pixel Buffer Region",
            "Programmed Behavior Response"
        ],
        "correctIndex": 1
    }
]
QUIZ_END -->
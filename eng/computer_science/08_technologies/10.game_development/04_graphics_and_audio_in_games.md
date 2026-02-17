# 🎨 Graphics and Audio in Games

## 📑 Contents
1. [Fundamentals of Game Graphics](#fundamentals-of-game-graphics)
2. [Types of Graphics in Games](#types-of-graphics-in-games)
3. [Rendering Techniques](#rendering-techniques)
4. [Audio in Games](#audio-in-games)
5. [Graphics and Audio Optimization](#graphics-and-audio-optimization)
6. [Tools and Formats](#tools-and-formats)

---

## 1. 🖼️ Fundamentals of Game Graphics

### 🧮 3D Graphics: From 3D to 2D
The process of transforming a 3D world into a 2D image on screen:

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    A[3D Model] --> B[Vertex Shader]
    B --> C[Tessellation]
    C --> D[Geometry Shader]
    D --> E[Fragment Shader]
    E --> F[Rasterization]
    F --> G[2D Image on Screen]
```

### 🧩 Core Components:
*   **Models** — 3D objects (meshes)
*   **Textures** — images for surfaces
*   **Shaders** — programs for graphics processing
*   **Lighting** — light sources and shadows
*   **Camera** — player's viewpoint

### 📐 Geometry Types:
*   **Polygons** — triangles, quads
*   **Splines** — curved lines
*   **NURBS** — mathematical surfaces

---

## 2. 🎯 Types of Graphics in Games

### 📷 3D Graphics
*   **Advantages:**
    *   Realism
    *   Freedom of movement
    *   Interactivity
*   **Disadvantages:**
    *   High hardware requirements
    *   Complexity in development
*   **Examples:** Assassin's Creed, The Witcher 3, Red Dead Redemption, **Fortnite**

**Example from Fortnite:**
In Fortnite, 3D graphics with unique stylized visual design are used, combining:
- Bright, saturated colors to create a comic book style
- Stylized character models with exaggerated features
- Dynamic building system with the ability to quickly change the landscape
- Particle effects for visualizing construction, destruction, and combat actions
- Lighting system adapted to time of day and weather conditions

### 🎨 2.5D (Isometric, Parallax Scrolling)
*   **Advantages:**
    *   Visually interesting
    *   Balance between 2D and 3D
    *   Efficiency
*   **Disadvantages:**
    *   Limited camera freedom
*   **Examples:** Diablo II, Papers, Please, Limbo

### 🖌️ 2D Graphics
*   **Advantages:**
    *   Lower hardware requirements
    *   Faster development
    *   Unique style
*   **Disadvantages:**
    *   Limited depth
*   **Examples:** Super Mario Bros, Cuphead, Celeste

### 🎞️ Pixel Art
*   **Advantages:**
    *   Nostalgia
    *   Small file sizes
    *   Unique aesthetic charm
*   **Disadvantages:**
    *   Requires artistic skills
    *   Time-intensive to create
*   **Examples:** Shovel Knight, Stardew Valley, Hyper Light Drifter

---

## 3. 🚀 Rendering Techniques

### 🌟 Lighting:
*   **Ambient Light** — general scene lighting
*   **Directional Light** — sun, parallel rays
*   **Point Light** — light bulb, emits in all directions
*   **Spot Light** — flashlight, cone of light

**Example from Fortnite:**
In Fortnite, a dynamic lighting system is used that:
- Changes depending on the time of day in the game world
- Creates realistic shadows for buildings and characters
- Highlights weapon effects and abilities using directional lighting
- Uses ambient lighting to create a bright, comic book atmosphere

### 🎨 Shading Techniques:
*   **Flat Shading** — simple shading (all polygons same color)
*   **Gouraud Shading** — smooth shading between vertices
*   **Phong Shading** — detailed surface shading

### 🌌 Advanced Effects:
*   **Normal Mapping** — simulate details on smooth surfaces
*   **Parallax Mapping** — texture depth
*   **Specular Mapping** — reflective properties
*   **Environment Mapping** — environmental reflections

**Example from Fortnite:**
In Fortnite, the following advanced graphics effects are applied:
- **Normal Mapping** — for adding detail to building surfaces and character models
- **Specular Mapping** — for creating realistic reflections on weapons and shiny items
- **Particle Effects** — for visualizing construction, destruction, shooting, and abilities
- **Dynamic Materials** — for changing appearance of materials during construction and destruction

### 🎞️ Post-Processing Techniques:
*   **Bloom** — bright objects glow
*   **Motion Blur** — blur from movement
*   **Depth of Field** — blur outside focus
*   **Anti-Aliasing** — smooth jagged edges

**Example from Fortnite:**
In Fortnite, the following post-processing techniques are used:
- **Bloom** — to create bright glow around weapons, abilities, and glowing elements
- **Motion Blur** — to enhance the feeling of speed during fast movement
- **Anti-Aliasing** — to smooth out stair-stepped lines on buildings and characters
- **Chromatic Aberration** — for effects when activating certain abilities

### 🌐 Global Illumination:
*   **Lightmapping** — pre-calculated lighting
*   **Real-time GI** — dynamic global illumination
*   **Ray Tracing** — physically accurate light modeling

**Example from Fortnite:**
In Fortnite, pre-calculated lighting (lightmapping) is used for:
- Creating realistic lighting on static map elements
- Optimizing performance on mobile devices and lower-end PCs
- Maintaining the game's bright, comic book atmosphere

---

## 4. 🎵 Audio in Games

### 🎼 Audio Components:
*   **Music** — soundtrack, atmosphere
*   **Sound Effects (SFX)** — footsteps, shots, explosions
*   **Voice Acting** — character speech
*   **Ambience** — background sounds (wind, rain)

**Example from Fortnite:**
In Fortnite, audio plays a key role in gameplay:
- **Music** — dynamic music that intensifies tension during combat
- **Sound Effects** — clear sounds of shooting, reloading, building, and destroying structures
- **Voice Acting** — character lines and voice notifications of events
- **Ambience** — nature sounds, wind, music in the safe zone

### 🎧 Spatial Audio:
*   **3D Audio** — sound from specific direction
*   **HRTF (Head-Related Transfer Function)** — sound perception modeling
*   **Dynamic Audio** — sound changes based on environment

**Example from Fortnite:**
In Fortnite, spatial audio is critically important for gameplay:
- **3D Audio** — allows players to determine enemy location by sounds of building, footsteps, and shooting
- **Dynamic Audio** — sounds change depending on whether the player is inside or outside buildings
- **Distance Attenuation** — sounds become quieter with distance, helping determine event remoteness

### 🎚️ Audio Systems:
*   **FMOD** — professional audio system
*   **Wwise** — audio engine by Audiokinetic
*   **Built-in Audio** — built-in systems (Unity, Unreal)

### 🎼 Audio Formats:
*   **WAV** — high quality, large size
*   **MP3** — lossy compression
*   **OGG** — lossless compression, good for games
*   **FLAC** — lossless compression

---

## 5. ⚡ Graphics and Audio Optimization

### 🖼️ Graphics Optimization:

#### LOD (Level of Detail)
*   Using less detailed models at a distance
*   Reducing polygon count
*   Reducing texture resolution

**Example from Fortnite:**
In Fortnite, various graphics optimization methods are used:
- **LOD** — simplified models of buildings and characters at a distance
- **Occlusion Culling** — buildings behind other objects are not rendered
- **Texture Streaming** — loading textures as needed, especially during fast movement
- **Batch Rendering** — combining similar objects (e.g., multiple identical buildings) for more efficient rendering

#### Occlusion Culling
*   Not rendering objects hidden by other objects
*   Increases performance

#### Frustum Culling
*   Not rendering objects outside camera's view
*   Saves GPU resources

#### Texture Streaming
*   Loading textures as needed
*   Reduces memory usage

#### Batch Rendering
*   Combining similar objects for rendering
*   Reduces number of draw calls

### 🎵 Audio Optimization:
*   **Compression** — compressing audio files
*   **Streaming** — loading audio as needed
*   **Occlusion** — reducing volume when obstructed
*   **Distance Attenuation** — reducing volume with distance

**Example from Fortnite:**
In Fortnite, the following audio optimization methods are applied:
- **Compression** — efficient compression of audio files to reduce download size
- **Streaming** — loading audio as needed, especially for music tracks
- **Occlusion** — reducing volume of sounds through walls and structures
- **Distance Attenuation** — sounds become quieter with distance, which helps in map orientation

---

## 6. 🛠️ Tools and Formats

### 🎨 Graphics Tools:
*   **Blender** — 3D modeling, animation (free)
*   **Maya/3ds Max** — professional 3D tools
*   **Photoshop/GIMP** — 2D graphics, textures
*   **Substance Painter** — 3D model texturing
*   **Aseprite** — pixel art

### 📁 Graphics Formats:
*   **FBX** — 3D models with animation
*   **OBJ** — simple 3D format
*   **PNG** — textures with transparency
*   **DDS** — textures with mipmaps
*   **TGA** — high-quality textures

### 🎵 Audio Tools:
*   **Audacity** — audio editing (free)
*   **FMOD Studio** — professional audio environment
*   **Wwise** — audio engine
*   **Reaper** — digital audio workstation

### 🎨 Visual Effects:
*   **Particle Systems** — smoke, fire, explosions
*   **Shaders** — visual effects (water, lava)
*   **Post-processing** — final image processing

---

## 7. 🚀 Conclusion

Graphics and audio are key elements that create atmosphere and immersion in the game world. Modern games use sophisticated rendering and audio processing techniques to create impressive visual and sound effects. However, developers must balance quality and performance to ensure the game is accessible on various devices.

The next materials will cover physics in games, network development, and other important aspects.

<!-- QUIZ_START
[
    {
        "question": "What does the abbreviation LOD mean in the context of game graphics?",
        "options": [
            "Lighting Over Distance",
            "Level of Detail",
            "Layered Object Display",
            "Linear Optimization Data"
        ],
        "correctIndex": 1
    },
    {
        "question": "Which audio format is typically used for lossless compression in games?",
        "options": [
            "MP3",
            "WAV",
            "OGG",
            "AAC"
        ],
        "correctIndex": 2
    },
    {
        "question": "What is Normal Mapping in the context of 3D graphics?",
        "options": [
            "A technique for smoothing edges",
            "Simulating details on smooth surfaces",
            "A method for lighting scenes",
            "A texture compression technique"
        ],
        "correctIndex": 1
    }
]
QUIZ_END -->
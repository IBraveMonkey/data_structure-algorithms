# ⚡ Physics in Games

## 📑 Contents
1. [Fundamentals of Game Physics](#fundamentals-of-game-physics)
2. [Types of Physical Simulations](#types-of-physical-simulations)
3. [Physics Engines](#physics-engines)
4. [Collisions and Collision Detection](#collisions-and-collision-detection)
5. [Realistic Physics vs Game Physics](#realistic-physics-vs-game-physics)
6. [Application of Physics in Game Design](#application-of-physics-in-game-design)

---

## 1. 🧮 Fundamentals of Game Physics

### 🌍 What is Game Physics?
**Game Physics** is the simulation of physical laws in a virtual environment to create realistic object behavior.

### 🔄 Basic Physical Concepts:
*   **Gravity** — force pulling objects toward the ground
*   **Inertia** — body maintains state of rest or motion
*   **Momentum** — mass times velocity
*   **Friction** — force opposing motion
*   **Elasticity** — ability of objects to recover from deformation

**Example from Fortnite:**
In Fortnite, physics plays a key role in gameplay:
- **Gravity** — affects jumps, falling from heights, and gliding with parachute
- **Momentum** — determines character movement during jumps and landings
- **Friction** — affects sliding and stopping of characters
- **Elasticity** — used in some utility items and abilities

### 📐 Basic Formulas:
*   **F = ma** (Force = mass × acceleration)
*   **v = u + at** (Velocity = initial velocity + acceleration × time)
*   **s = ut + ½at²** (Distance = initial velocity × time + ½ × acceleration × time²)

---

## 2. 🧩 Types of Physical Simulations

### 🏃‍♂️ Kinematics (motion):
*   **Position** — coordinates of object
*   **Velocity** — change in position over time
*   **Acceleration** — change in velocity over time

**Example from Fortnite:**
In Fortnite, kinematics is used for:
- **Position** — tracking player and enemy locations on the map
- **Velocity** — determining movement speed during walking, running, and jumping
- **Acceleration** — modeling acceleration when starting movement and deceleration when landing

### 🏗️ Rigid Body Dynamics:
*   **Mass** — measure of inertia
*   **Force** — applied to object
*   **Momentum** — mass × velocity
*   **Rotation** — torque, angular velocity

**Example from Fortnite:**
In Fortnite, rigid body dynamics are applied for:
- **Mass** — determining weight of character and items
- **Force** — modeling recoil from shots and explosions
- **Momentum** — calculating movement during jumps and taking damage
- **Rotation** — weapon animation and camera movement

### 🌊 Fluid Simulation:
*   **Water** — waves, currents
*   **Gases** — smoke, clouds
*   **Granular materials** — sand, snow

**Example from Fortnite:**
In Fortnite, fluid simulation is used for:
- **Water** — visual effects when contacting water
- **Gases** — fog and smoke effects from abilities and explosions
- **Granular materials** — effects when destroying buildings and collecting resources

### 🧊 Deformations and Destruction:
*   **Elastic bodies** — return to original shape
*   **Plastic bodies** — retain deformation
*   **Breakable objects** — break under stress

**Example from Fortnite:**
In Fortnite, destruction physics is a key element of gameplay:
- **Breakable objects** — walls, floors, and stairs can be destroyed by bullets and explosions
- **Dynamic destructions** — buildings are partially destroyed, depending on the hit location
- **Visual effects** — destructions are accompanied by particles and sounds

### 🎾 Specific Effects:
*   **Ricochet** — bouncing off surfaces
*   **Sliding** — movement along surfaces
*   **Rolling** — movement with rotation

**Example from Fortnite:**
In Fortnite, specific physical effects include:
- **Ricochet** — bullets can bounce off building surfaces
- **Sliding** — players can slide when descending slopes and inclined surfaces
- **Rolling** — effects when using certain items and abilities

---

## 3. 🏗️ Physics Engines

### 🧰 Popular Physics Engines:

#### Box2D
*   **Type:** 2D physics
*   **Language:** C++
*   **Application:** 2D games, platformers
*   **Features:**
    *   High accuracy
    *   Stability
    *   Easy integration
*   **Example Games:** Angry Birds, Limbo

#### Bullet Physics
*   **Type:** 3D physics
*   **Language:** C++
*   **Application:** 3D games, simulators
*   **Features:**
    *   Soft body support
    *   Complex collisions
    *   High performance
*   **Example Games:** Grand Theft Auto, many AAA titles

#### PhysX (NVIDIA)
*   **Type:** 3D physics
*   **Language:** C++
*   **Application:** AAA games
*   **Features:**
    *   GPU acceleration
    *   Realistic simulations
    *   Engine integration
*   **Example Games:** Assassin's Creed, Call of Duty, **Fortnite**

**Example from Fortnite:**
Fortnite uses PhysX for:
- Building destruction simulation
- Character and item physics movement
- Environmental interaction
- Particle effects and dynamic objects

#### Havok
*   **Type:** 3D physics
*   **Language:** C++
*   **Application:** AAA games
*   **Features:**
    *   High precision
    *   Professional tools
    *   Licensing
*   **Example Games:** The Elder Scrolls, Fallout

#### Chipmunk
*   **Type:** 2D physics
*   **Language:** C
*   **Application:** 2D mobile games
*   **Features:**
    *   Lightweight
    *   Fast performance
    *   Ease of use
*   **Example Games:** Tiny Wings, many mobile games

---

## 4. 🧊 Collisions and Collision Detection

### 🎯 Collision Types:
*   **Static** — immovable objects (walls, floor)
*   **Kinematic** — moving along predetermined path
*   **Dynamic** — objects subject to physics

**Example from Fortnite:**
In Fortnite, different types of collisions are used:
- **Static** — built walls, floors, and stairs (until destroyed)
- **Kinematic** — moving platforms and some animated objects
- **Dynamic** — characters, projectiles, destroyed parts of buildings

### 📐 Collision Shapes:
*   **Box** — simple rectangular shapes
*   **Circle/Sphere** — circles and spheres
*   **Capsule** — capsules (cylinder with hemispheres)
*   **Polygon/Mesh** — complex polygonal shapes
*   **Terrain** — complex landscapes

**Example from Fortnite:**
In Fortnite, various collision shapes are used:
- **Box** — for buildings (walls, floors, stairs)
- **Capsule** — for characters and weapons
- **Sphere** — for projectiles and some effects
- **Terrain** — for the base of the game world and landscape

### 🔍 Detection Algorithms:
*   **Bounding Volume Hierarchy (BVH)** — hierarchy of bounding volumes
*   **Spatial Hashing** — dividing space into cells
*   **Sweep and Prune** — axis-based sorting
*   **GJK Algorithm** — for convex shapes

**Example from Fortnite:**
In Fortnite, collision detection algorithms are used for:
- **BVH** — quick collision checking between numerous buildings
- **Spatial Hashing** — efficient collision checking in densely built areas
- **Sweep and Prune** — optimization of collision checks for moving objects (characters, projectiles)

### 🔄 Collision Response:
*   **Bounce** — elastic collision
*   **Stick** — objects adhere to each other
*   **Penetration** — objects pass through each other
*   **Destruction** — objects break upon collision

**Example from Fortnite:**
In Fortnite, collision responses include:
- **Bounce** — projectiles can bounce off building surfaces
- **Destruction** — buildings are destroyed when hit by projectiles
- **Movement restriction** — characters collide with buildings and don't pass through them

---

## 5. 🎮 Realistic Physics vs Game Physics

### 📊 Comparison:

| Aspect | Realistic Physics | Game Physics |
|--------|-------------------|--------------|
| **Accuracy** | Maximum | Acceptable |
| **Performance** | Low | High |
| **Predictability** | Yes | Yes, but may be altered |
| **Playability** | May be poor | Priority |
| **Goal** | Physical accuracy | Player satisfaction |

### 🎯 Game Physics:
*   **Simplifications** — ignoring certain physical laws
*   **Intuitiveness** — behavior should be understandable to player
*   **Predictability** — player should be able to predict outcomes
*   **Fun** — physics can be "unrealistic" but fun

### 🚀 Examples of Game Physics:
*   **Super Mario** — jumps higher than in reality
*   **Tony Hawk's Pro Skater** — impossible tricks
*   **Angry Birds** — simplified ballistics
*   **Portal** — violation of physical laws
*   **Fortnite** — simplified ballistics and building physics

**Example from Fortnite:**
Fortnite demonstrates game physics where:
- Jumps and movement are optimized for convenience and speed
- Projectile ballistics are simplified for better gameplay
- Building physics are accelerated for dynamic gameplay
- Gravity and movement are balanced for fun and competitive gameplay

---

## 6. 🎨 Application of Physics in Game Design

### 🧩 Physics-Based Mechanics:
*   **Jumping** — overcoming obstacles
*   **Ballistics** — shooting, throwing
*   **Destruction** — satisfaction from destruction
*   **Interaction** — pushing, lifting, rolling
*   **Transportation** — driving, flying, swimming

**Example from Fortnite:**
In Fortnite, physics is used in the following mechanics:
- **Jumping** — key element of movement and tactics
- **Ballistics** — shooting from various weapons
- **Destruction** — core gameplay element, ability to destroy buildings
- **Interaction** — resource collection, building
- **Flying** — with parachute from the island at the start of the match

### 🎯 Examples of Use:
*   **Angry Birds** — entire game based on ballistics
*   **Half-Life** — physics used for realism
*   **Portal** — physics used for puzzles
*   **Teardown** — physics as core gameplay
*   **Fortnite** — physics as core of dynamic building and destruction

**Example from Fortnite:**
In Fortnite, physics is used for:
- **Building** — instant creation of structures
- **Destructions** — dynamic destruction of buildings
- **Movement** — jumping, running, parachuting
- **Combat system** — ballistics, damage, recoil

### 🧠 Psychological Aspect:
*   **Satisfaction** — destruction, jumping, movement
*   **Intuition** — player feels the "weight" of objects
*   **Immersion** — physics enhances immersion
*   **Learning** — player learns the game's physical rules

### 🎮 Balance:
*   **Complexity** — overly complex physics can confuse
*   **Predictability** — player should understand consequences
*   **Fun** — physics should enhance enjoyment
*   **Optimization** — physics shouldn't slow down the game

---

## 7. 🚀 Conclusion

Physics in games is not just a copy of the real world, but a tool for creating engaging gameplay. Well-designed game physics makes the world more alive and predictable for the player, while allowing designers to create unique mechanics and experiences.

The next materials will cover network game development, monetization, and other important aspects.

<!-- QUIZ_START
[
    {
        "question": "Which physics engine is primarily used for 2D games?",
        "options": [
            "Bullet Physics",
            "PhysX",
            "Box2D",
            "Havok"
        ],
        "correctIndex": 2
    },
    {
        "question": "What does the abbreviation BVH mean in the context of collision detection?",
        "options": [
            "Basic Vector Helper",
            "Bounding Volume Hierarchy",
            "Binary Vertex Hull",
            "Balanced Velocity Handler"
        ],
        "correctIndex": 1
    },
    {
        "question": "Which formula describes Newton's second law?",
        "options": [
            "E = mc²",
            "F = ma",
            "v = d/t",
            "p = mv"
        ],
        "correctIndex": 1
    }
]
QUIZ_END -->
<template>
  <div id="effect-square" class="effect-square" />
</template>

<script>
import "particles.js"
export default {
  name: 'EffectSquare',
  mounted () {
    this.initParticles()
  },
  methods: {
    initParticles () {
      window.particlesJS('effect-square',
        {
          particles: {
            number: {
              value: 25,
              density: {
                enable: true,
                value_area: 1000
              }
            },
            color: {
              value: ['#7fffbf', '#7fff7f']
            },

            shape: {
              type: 'edge',
              stroke: {
                width: 0,
                color: '#000000'
              },
              polygon: {
                nb_sides: 5
              },
              image: {
                src: 'img/github.svg',
                width: 100,
                height: 100
              }
            },

            opacity: {
              value: 0.5,
              random: true,

              anim: {
                enable: false,
                speed: 1,
                opacity_min: 0.1,
                sync: false
              }
            },

            size: {
              value: 25,
              random: true,
              anim: {
                enable: true,
                speed: 20,
                size_min: 5,
                sync: true
              }
            },
            line_linked: {
              enable: false,
              distance: 150,
              color: '#ffffff',
              opacity: 0.4,
              width: 1
            },
            move: {
              enable: true,
              speed: 0,
              direction: 'none',
              random: false,
              straight: false,
              out_mode: 'out',
              attract: {
                enable: true,
                rotateX: 1000,
                rotateY: 1000
              }
            }
          },

          interactivity: {
            detect_on: 'canvas',
            events: {
              onhover: {
                enable: false,
                mode: 'bubble'
              },
              onclick: {
                enable: false,
                mode: 'push'
              },
              resize: true
            },
            modes: {
              grab: {
                distance: 300,
                line_linked: {
                  opacity: 1
                }
              },
              bubble: {
                distance: 100,
                size: 7.5,
                duration: 2,
                opacity: 8,
                speed: 3
              },
              repulse: {
                distance: 1
              },
              push: {
                particles_nb: 4
              },
              remove: {
                particles_nb: 2
              }
            }
          },
          retina_detect: true
        }
      )
    }
  }
}
</script>

<style lang="scss" scoped>
#effect-square{
  position: fixed;
  z-index: -1;
  width: 100%;
  height: 100%;
}
</style>

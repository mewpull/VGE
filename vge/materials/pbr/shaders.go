//

package pbr

import _ "embed"

//go:generate glslangValidator -V pbr.vert.glsl -o pbr.vert.spv
//go:generate glslangValidator -V -DSKINNED=1 pbr.vert.glsl -o pbr.vert_skin.spv
//go:generate glslangValidator -V pbr.frag.glsl -o pbr.frag.spv

//go:embed pbr.vert.spv
var pbr_vert_spv []byte

//go:embed pbr.frag.spv
var pbr_frag_spv []byte

//go:embed pbr.vert_skin.spv
var pbr_vert_skin_spv []byte

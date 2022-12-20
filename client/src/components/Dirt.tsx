import { DirtProps } from "../interfaces";

/**
 * A div to be nested inside a Box
 * 
 * 125px wide, 20 px tall
 * 
 * To be used when clicking to clone flowers. (Destination Box will be the box associated with this Dirt.)
 * 
 * Styling is done inside Box scss since Dirts are dumb componenents.
 * 
 * @param props Props. See DirtProps documentation
 * @return Colored div 
 */
function Dirt(props: DirtProps) {
    return (
        <div className="dirt"></div>
    )
}

export default Dirt;
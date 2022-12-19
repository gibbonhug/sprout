/**
 * See Go 'Flower' struct
 * - id: Unique unsigned ID
 * - colorPetal: Hex color of petals (without #)
*/
export interface FlowerProps {
    id: number;
    colorPetal: string;
}

/**
 * See Go 'Box' struct
 * - id: Unique unsigned ID
 */
export interface BoxProps {
    id: number;
}
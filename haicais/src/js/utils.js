export const getRandom = (arr) => arr[Math.floor(Math.random() * arr.length)]

export function filterTheme(arr = [], theme){
    const adj = arr.filter(item => item.themes.some(i => i === theme))
    return adj
}
declare interface ModalityRouteMeta {
  title: string
  searchable?: boolean
  menu?: {
    name: string
    group: string
    icon?: unknown
    priority?: number
  }
}

declare interface ModalityPluginDefinition {
  components: Record<string, unknown>
  routes: Array<{
    parentName: string
    route: {
      path: string
      name: string
      component: unknown
      meta: ModalityRouteMeta
    }
  }>
  extensionPoints: Record<string, unknown>
}

declare function defineModalityPlugin(
  definition: ModalityPluginDefinition,
): ModalityPluginDefinition

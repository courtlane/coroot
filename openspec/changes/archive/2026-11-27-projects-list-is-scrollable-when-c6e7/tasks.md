## 1. Поиск в меню проектов (`front/src/App.vue`)
- [x] 1.1 Добавить в `data()` поля `projectsMenu: false` и `projectSearch: ''`, в `computed` — `filteredProjects` (регистронезависимый поиск подстроки в `p.name`, при пустом/`null` запросе — `projects` без изменений) и `showProjectSearch` (`projects.length > 7`).
- [x] 1.2 Повесить `v-model="projectsMenu"` на `v-menu` проектов; в `watch` при `projectsMenu === false` сбрасывать `projectSearch`, при `true` — через `$nextTick` ставить фокус в поле поиска (через `ref`), если оно показано.
- [x] 1.3 Над `<div class="projects">` добавить обёртку с `v-if="showProjectSearch"`, `@click.stop` и `@keydown` (останавливать всплытие для всех клавиш, кроме Esc), внутри — `v-text-field` (`v-model="projectSearch"`, `dense`, `outlined` или `solo-inverted`, `hide-details`, `clearable`, `prepend-inner-icon="mdi-magnify"`, `placeholder="search"`, `ref`).
- [x] 1.4 Заменить `v-for="p in projects"` на `v-for="p in filteredProjects"`; при непустом запросе и пустом результате показать неактивный пункт «no matching projects» внутри блока `.projects`.
- [x] 1.5 Добавить в `<style scoped>` минимальные стили для обёртки поля (отступы, ширина как у меню), проверить вид в тёмном меню.

## 2. Проверка
- [ ] 2.1 Тесты: автотестов для фронтенда в проекте нет; выполнить `cd front && npm run lint` и `go test ./...` (бэкенд не меняется, должен оставаться зелёным).
- [ ] 2.2 Ручная проверка (собрать фронт `npm run build-dev`): ≤7 проектов — поля нет, меню как раньше; >7 — поле есть, фокус в нём при открытии, ввод «test1» оставляет только совпадения, регистр не важен, «no matching projects» при отсутствии совпадений.
- [ ] 2.3 Ручная проверка: ввод и клик в поле не закрывают меню, Esc закрывает; клик по найденному проекту открывает его; после повторного открытия запрос пуст; «new project» виден для не-readonly; Ctrl/Cmd+K открывает «Go to…»; проверить на маленьком окне.

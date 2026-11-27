## 1. Разметка меню проектов
- [x] 1.1 В `front/src/App.vue` убрать `max-height="80vh"` у `<v-menu v-if="user" dark right offset-x tile ...>` меню «Project»
- [x] 1.2 Обернуть `<v-list-item v-for="p in projects" ...>` в `<div class="projects">`; пункты «new project» и «no projects available» оставить после этого `div`
- [x] 1.3 Добавить в `<style scoped>` правило `.projects { max-height: 70vh; overflow-y: auto; }`

## 2. Проверка
- [ ] 2.1 Выполнить `cd front && npm run lint` и `npm run fmt` — без ошибок и лишних изменений
- [ ] 2.2 Автотестов фронта в проекте нет: собрать фронт (`npm run build-dev` или `npm run build`), запустить Coroot со свежей статикой (`DEVELOPER_MODE=true go run main.go` или пересобрать бинарник) и вручную проверить с ~30 проектами: список проектов прокручивается до последнего, «new project» виден внизу меню без прокрутки, меню не выходит за край окна (в том числе на маленьком окне высотой ~500px)
- [ ] 2.3 Вручную проверить с 2–3 проектами: меню без полосы прокрутки и выглядит как раньше; проверить readonly-пользователя (нет «new project»)
- [ ] 2.4 Выполнить `go test ./...` — бэкенд не изменён, тесты проходят

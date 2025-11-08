-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.youtube
ADD COLUMN track_id UUID REFERENCES public.tracks(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.youtube
DROP COLUMN track_id;
-- +goose StatementEnd

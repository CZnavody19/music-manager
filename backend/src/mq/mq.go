package mq

import (
	"context"
	"encoding/json"

	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/internal/discord"
	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/plex"
	"github.com/CZnavody19/music-manager/src/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type MessageQueue struct {
	conn         *amqp.Connection
	discord      *discord.Discord
	plex         *plex.Plex
	musicbrainz  *musicbrainz.MusicBrainz
	success_chan <-chan amqp.Delivery
	fail_chan    <-chan amqp.Delivery
}

func NewMessageQueue(conn *amqp.Connection, discord *discord.Discord, plex *plex.Plex, mb *musicbrainz.MusicBrainz) (*MessageQueue, error) {
	ctx := context.Background()

	chann, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = chann.QueueDeclare("downloads_complete.success", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_, err = chann.QueueDeclare("downloads_complete.fail", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = chann.QueueBind("downloads_complete.success", "success", "downloads_complete", false, nil)
	if err != nil {
		return nil, err
	}

	err = chann.QueueBind("downloads_complete.fail", "fail", "downloads_complete", false, nil)
	if err != nil {
		return nil, err
	}

	success_chan, err := chann.ConsumeWithContext(ctx, "downloads_complete.success", "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	fail_chan, err := chann.ConsumeWithContext(ctx, "downloads_complete.fail", "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	mq := &MessageQueue{
		conn:         conn,
		discord:      discord,
		plex:         plex,
		musicbrainz:  mb,
		success_chan: success_chan,
		fail_chan:    fail_chan,
	}

	go mq.failWorker(ctx)
	go mq.successWorker(ctx)

	return mq, nil
}

func (mq *MessageQueue) Reload(ctx context.Context, service string) error {
	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "reload", service, false, false, amqp.Publishing{})

	if err != nil {
		return err
	}

	return nil
}

func (mq *MessageQueue) Download(ctx context.Context, track *domain.Track, service string) error {
	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}

	body, err := json.Marshal(track)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "downloads", service, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return err
	}

	return nil
}

func (mq *MessageQueue) successWorker(ctx context.Context) {
	zap.S().Info("Starting success worker")

	for message := range mq.success_chan {
		zap.S().Info("Received success message")

		var msg domain.CompleteMessage

		err := json.Unmarshal(message.Body, &msg)
		if err != nil {
			zap.S().Error("Failed to unmarshal success message", err)
			continue
		}

		err = mq.musicbrainz.MarkDownloaded(ctx, msg.Track.ID, true)
		if err != nil {
			zap.S().Error("Failed to mark track as downloaded", err)
		}

		err = mq.plex.RefreshLibrary(ctx)
		if err != nil {
			zap.S().Error("Failed to refresh plex library", err)
		}

		zap.S().Info("Plex library refreshed")

		err = mq.discord.SendMessage(ctx, &domain.DiscordMessage{
			Title: "Download Complete",
			Color: utils.IntPtr(65280), // green
			Fields: []domain.DiscordMessageField{
				{
					Name:   "Title",
					Value:  msg.Track.Title,
					Inline: false,
				},
				{
					Name:   "Artist",
					Value:  msg.Track.Artist,
					Inline: false,
				},
			},
		})
		if err != nil {
			zap.S().Error("Failed to send discord message", err)
			return
		}

		zap.S().Info("Discord message sent")
	}
}

func (mq *MessageQueue) failWorker(ctx context.Context) {
	zap.S().Info("Starting fail worker")

	for message := range mq.fail_chan {
		zap.S().Info("Received fail message")

		var msg domain.CompleteMessage

		err := json.Unmarshal(message.Body, &msg)
		if err != nil {
			zap.S().Error("Failed to unmarshal success message", err)
			continue
		}

		err = mq.musicbrainz.MarkDownloaded(ctx, msg.Track.ID, false)
		if err != nil {
			zap.S().Error("Failed to mark track as not downloaded", err)
		}

		err = mq.discord.SendMessage(ctx, &domain.DiscordMessage{
			Title:       "Download Failed",
			Description: msg.Error,
			Color:       utils.IntPtr(16711680), // red
			Fields: []domain.DiscordMessageField{
				{
					Name:   "Title",
					Value:  msg.Track.Title,
					Inline: false,
				},
				{
					Name:   "Artist",
					Value:  msg.Track.Artist,
					Inline: false,
				},
			},
		})
		if err != nil {
			zap.S().Error("Failed to send discord message", err)
			return
		}

		zap.S().Info("Discord message sent")
	}
}

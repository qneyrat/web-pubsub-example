<?php
declare(strict_types=1);

namespace App\EventSubscriber;

use App\Broker\MessagePublisher;
use App\Event\MessageEvent;
use Symfony\Component\EventDispatcher\EventSubscriberInterface;

class MessageCreatedSubscriber implements EventSubscriberInterface
{
    /**
     * @var MessagePublisher
     */
    private $publisher;

    /**
     * MessageCreatedSubscriber constructor.
     * @param MessagePublisher $publisher
     */
    public function __construct(MessagePublisher $publisher)
    {
        $this->publisher = $publisher;
    }

    /**
     * {@inheritdoc}
     */
    public static function getSubscribedEvents()
    {
        return [
            MessageEvent::CREATED => 'onMessageCreated'
        ];
    }

    public function onMessageCreated(MessageEvent $event)
    {
        $this->publisher->messageAdded($event->getMessage()->getConversation(), $event->getMessage());
    }
}

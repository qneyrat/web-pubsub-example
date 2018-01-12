<?php
declare(strict_types=1);

namespace App\EventSubscriber\Doctrine;

use App\Document\Conversation;
use App\Event\ConversationEvent;
use Doctrine\Common\EventSubscriber;
use Doctrine\ODM\MongoDB\Event\LifecycleEventArgs;
use Symfony\Component\Cache\Simple\ArrayCache;
use Symfony\Component\EventDispatcher\EventDispatcherInterface;

class ConversationSubscriber implements EventSubscriber
{
    /**
     * @var ArrayCache
     */
    private $cache;

    /**
     * @var EventDispatcherInterface
     */
    private $eventDispatcher;

    /**
     * ConversationSubscriber constructor.
     * @param ArrayCache $cache
     * @param EventDispatcherInterface $eventDispatcher
     */
    public function __construct(ArrayCache $cache, EventDispatcherInterface $eventDispatcher)
    {
        $this->cache = $cache;
        $this->eventDispatcher = $eventDispatcher;
    }

    /**
     * {@inheritdoc}
     */
    public function getSubscribedEvents()
    {
        return [
            'preUpdate',
            'postUpdate',
        ];
    }

    public function preUpdate(LifecycleEventArgs $args)
    {
        if (!$args->getDocument() instanceof Conversation) {
            return;
        }

        /** @var Conversation $conversation */
        $conversation = $args->getDocument();
        $newMessages = $conversation->getMessages()->getInsertedDocuments();
        $this->cache->set(spl_object_hash($conversation).'-messages-inserted', $newMessages);
    }

    public function postUpdate(LifecycleEventArgs $args)
    {
        if (!$args->getDocument() instanceof Conversation) {
            return;
        }

        /** @var Conversation $conversation */
        $conversation = $args->getDocument();
        $key = spl_object_hash($conversation).'-messages-inserted';
        $insertedMessages = $this->cache->get($key);
        if (!$insertedMessages) {
            return;
        }

        foreach ($insertedMessages as $message) {
            $this->eventDispatcher->dispatch(
              ConversationEvent::MESSAGE_CREATED,
              new ConversationEvent($conversation, $message)
            );
        }

        $this->cache->delete($key);
    }
}
